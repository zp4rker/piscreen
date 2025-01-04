from epd.epd2in13_V3 import EPD
from epd.gt1151 import GT1151, GT_Development
from threading import Condition, Thread
from PIL import Image, ImageDraw, ImageFont
import logging
from queue import Queue, Empty
import RPi.GPIO as GPIO
import signal


class PiScreen(object):
    epd = EPD()
    gt = GT1151()
    gt_dev = GT_Development()
    gt_dev_old = GT_Development()

    running = True
    busy = False
    current_screen = None

    touched = Condition()
    touch_thread = None
    pressed = False
    dragging = False
    event_queue = Queue()

    font = ImageFont.truetype("font.ttf", 14)


    def main(self):
        self.epd.init(self.epd.FULL_UPDATE)
        self.gt.GT_Init()
        self.gt.GT_Reset()
        GPIO.add_event_detect(self.gt.INT, GPIO.BOTH, callback=self.on_gpio_change)
        signal.signal(signal.SIGINT, self.on_ctrl_c)


        image = Image.new("1", (self.epd.height, self.epd.width), 255)
        draw = ImageDraw.Draw(image)

        draw.rectangle([(0,0),(self.epd.height, self.epd.width)], fill = 1)

        draw.text((125 - 2, 61), "Assalaamu", font = self.font, fill = 0, anchor = "rm")

        draw.rectangle([125,0, self.epd.height, 125], fill = 0)
        draw.text((125 + 2, 61), "'alaykum!", font = self.font, fill = 1, anchor = "lm")

        self.epd.display(self.epd.getbuffer(image))

        self.touch_thread = Thread(target=self.touch_loop)
        self.touch_thread.start()
        while(self.running):
                    try:
                        evt = self.event_queue.get(timeout=1)
                        logging.debug("Got event")
                        if not evt["dragging"]:
                            logging.info("(%s,%s)", evt["x"], evt["y"])
                        # current_screen.on_tap(250-evt['y'], evt['x'])
                    except Empty:
                        pass

        self.touch_thread.join()
        self.epd.sleep()
        self.epd.Dev_exit()


    def on_gpio_change(self, channel) :
            logging.debug("Edge detected")
            if(self.gt.digital_read(self.gt.INT) == 0) :
                with self.touched:
                    self.gt_dev.Touch = 1
                    logging.debug("Signalling")
                    self.touched.notify()
            else :
                logging.debug("Zero")
                self.gt_dev.Touch = 0


    def touch_loop(self):
            try:
                while self.running:
                    with self.touched:
                        while self.gt_dev.Touch != 1 and self.running:
                            self.touched.wait(timeout=1)
                        logging.debug("Scan")
                        self.gt.GT_Scan(self.gt_dev, self.gt_dev_old)

                        if self.gt_dev.TouchCount == 0:
                            if self.pressed:
                                logging.debug("Lifted: cnt: %d, flag: %d, x: %d, y: %d, s: %d" % (self.gt_dev.TouchCount, self.gt_dev.TouchpointFlag, self.gt_dev.X[0], self.gt_dev.Y[0], self.gt_dev.S[0]))
                                if not self.dragging:
                                    event = {
                                            "x": self.gt_dev.X[0],
                                            "y": self.epd.width - self.gt_dev.Y[0],
                                            "s": self.gt_dev.S[0],
                                            "dragging": self.dragging
                                    }
                                    logging.debug("Queuing event")
                                    if self.busy:
                                        logging.debug("Dropping touch event while busy")
                                    else:
                                        self.event_queue.put(event)
                                
                            else:
                                logging.debug("Unexpected: cnt: %d, flag: %d, x: %d, y: %d, s: %d" % (self.gt_dev.TouchCount, self.gt_dev.TouchpointFlag, self.gt_dev.X[0], self.gt_dev.Y[0], self.gt_dev.S[0]))
                            self.dragging = False
                            self.pressed = False
                        else:
                            if not self.pressed:
                                logging.debug("Pressed: cnt: %d, flag: %d, x: %d, y: %d, s: %d" % (self.gt_dev.TouchCount, self.gt_dev.TouchpointFlag, self.gt_dev.X[0], self.gt_dev.Y[0], self.gt_dev.S[0]))
                            else:
                                if(self.gt_dev.distance(self.gt_dev_old) > 3):
                                    self.dragging = True
                                    logging.debug("Dragged: cnt: %d, flag: %d, x: %d, y: %d, s: %d" % (self.gt_dev.TouchCount, self.gt_dev.TouchpointFlag, self.gt_dev.X[0], self.gt_dev.Y[0], self.gt_dev.S[0]))
                            self.pressed = True
            except Exception as e:
                logging.error(e)
                self.running = False

    
    def on_ctrl_c(self, sig, frame):
        logging.info("Shutting down...")
        self.running = False


logging.basicConfig(level=logging.INFO)
app = PiScreen()
app.main()