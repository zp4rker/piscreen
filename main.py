import logging
import signal
import time
from queue import Empty, Queue

import RPi.GPIO as GPIO
from PIL import Image, ImageDraw, ImageFont

from core.touch_listener import TouchListener
from epd.epd2in13_V3 import EPD
from epd.gt1151 import GT1151
from screen.clock import ClockScreen
from screen.test import TestScreen


class PiScreen(object):
    epd = EPD()
    gt = GT1151()
    old_image = None

    running = True
    busy = False
    current_screen = None
    prev_screen = None

    touch_thread = None
    event_queue = Queue()


    def main(self):
        self.setup()

        self.screens = {
            "test": TestScreen(self),
            "clock": ClockScreen(self)
        }

        self.current_screen = TestScreen(self)
        self.current_screen.start()

        while(self.running):
            try:
                evt = self.event_queue.get(timeout=1)
                if self.current_screen:
                    self.current_screen.accept_event(evt)
                logging.debug("Got event")
                if evt["swipe"]:
                    logging.debug("(%s,%s) -> (%s,%s) %s swipe %ss", evt["x1"], evt["y1"], evt["x2"], evt["y2"], evt["direction"], evt["duration"])
                else:
                    logging.debug("(%s,%s) %ss", evt["x"], evt["y"], evt["duration"])
            except Empty:
                pass

        self.cleanup()


    def draw(self, image, partial=True):
        buffer = self.epd.getbuffer(image)
        if not self.old_image:
            self.epd.init(self.epd.FULL_UPDATE)
            self.epd.displayPartBaseImage(buffer)
            self.epd.init(self.epd.PART_UPDATE)
            self.old_image = image
            return

        if image == self.old_image:
            return
        
        if partial:
            self.epd.displayPartial_Wait(buffer)
        else:
            self.epd.init(self.epd.FULL_UPDATE)
            self.epd.displayPartBaseImage(buffer)
            self.epd.init(self.epd.PART_UPDATE)

        self.old_image = image


    def change_screen(self, screen):
        if screen not in self.screens:
            return
        
        self.prev_screen = self.current_screen
        self.current_screen = self.screens[screen]
        self.current_screen.start()


    def back_screen(self):
        if not self.prev_screen:
            return
        
        self.current_screen = self.prev_screen
        self.prev_screen = None
        self.current_screen.start()


    def setup(self):
        signal.signal(signal.SIGINT, self.quit)

        self.epd.init(self.epd.FULL_UPDATE)
        self.gt.GT_Init()
        self.gt.GT_Reset()

        self.touch_thread = TouchListener(self)
        GPIO.add_event_detect(self.gt.INT, GPIO.BOTH, callback=self.touch_thread.on_gpio_change)
        self.touch_thread.start()


    def cleanup(self):
        self.epd.sleep()
        time.sleep(3)
        self.touch_thread.join()
        self.epd.Dev_exit()

    
    def quit(self, sig=None, frame=None):
        logging.info("Shutting down...")
        self.running = False



logging.basicConfig(level=logging.INFO)
app = PiScreen()
app.main()