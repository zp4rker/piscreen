from threading import Thread, Condition
import logging
from datetime import datetime
from epd.gt1151 import GT_Development

class TouchListener(Thread):
    gt_dev = GT_Development()
    gt_dev_old = GT_Development()

    touch_start = None
    first_touch = None

    touched = Condition()
    pressed = False
    dragging = False
    
    def __init__(self, app):
        super().__init__()
        self.app = app

    def run(self):
        try:
            while self.app.running:
                with self.touched:
                    while self.gt_dev.Touch != 1 and self.app.running:
                        self.touched.wait(timeout=1)
                    logging.debug("Scan")
                    self.app.gt.GT_Scan(self.gt_dev, self.gt_dev_old)

                    if self.gt_dev.TouchCount == 0:
                        if self.pressed:
                            touch_end = datetime.now()
                            logging.debug("Lifted: cnt: %d, flag: %d, x: %d, y: %d, s: %d" % (self.gt_dev.TouchCount, self.gt_dev.TouchpointFlag, self.gt_dev.X[0], self.gt_dev.Y[0], self.gt_dev.S[0]))
                            event = {
                                "x1": self.first_touch["x"],
                                "y1": self.first_touch["y"],
                                "s1": self.first_touch["s"],
                                "x2": self.gt_dev.X[0],
                                "y2": self.app.epd.width - self.gt_dev.Y[0],
                                "s2": self.gt_dev.S[0]
                            } if self.dragging else {
                                "x": self.gt_dev.X[0],
                                "y": self.app.epd.width - self.gt_dev.Y[0],
                                "s": self.gt_dev.S[0]
                            }
                            event["swipe"] = self.dragging
                            event["duration"] = (touch_end - self.touch_start).total_seconds()
                            if self.dragging:
                                event["direction"] = self.swipe_direction(event)
                            logging.debug("Queuing event")
                            if self.app.busy:
                                logging.debug("Dropping touch event while busy")
                            else:
                                self.app.event_queue.put(event)
                            
                            
                        else:
                            logging.debug("Unexpected: cnt: %d, flag: %d, x: %d, y: %d, s: %d" % (self.gt_dev.TouchCount, self.gt_dev.TouchpointFlag, self.gt_dev.X[0], self.gt_dev.Y[0], self.gt_dev.S[0]))
                        self.pressed = False
                        self.first_touch = None
                        self.touch_start = None
                        self.dragging = False
                    else:
                        if not self.pressed:
                            logging.debug("Pressed: cnt: %d, flag: %d, x: %d, y: %d, s: %d" % (self.gt_dev.TouchCount, self.gt_dev.TouchpointFlag, self.gt_dev.X[0], self.gt_dev.Y[0], self.gt_dev.S[0]))
                        else:
                            if(self.gt_dev.distance(self.gt_dev_old) > 1):
                                self.dragging = True
                                logging.debug("Dragged: cnt: %d, flag: %d, x: %d, y: %d, s: %d" % (self.gt_dev.TouchCount, self.gt_dev.TouchpointFlag, self.gt_dev.X[0], self.gt_dev.Y[0], self.gt_dev.S[0]))
                        self.pressed = True
                        if not self.first_touch:
                            self.first_touch = {
                                "x": self.gt_dev.X[0],
                                "y": self.gt_dev.Y[0],
                                "s": self.gt_dev.S[0]
                            }
                        if not self.touch_start:
                            self.touch_start = datetime.now()
        except Exception as e:
            logging.error(e)
            self.app.running = False


    def on_gpio_change(self, channel) :
            logging.debug("Edge detected")
            if(self.app.gt.digital_read(self.app.gt.INT) == 0) :
                with self.touched:
                    self.gt_dev.Touch = 1
                    logging.debug("Signalling")
                    self.touched.notify()
            else:
                logging.debug("Zero")
                self.gt_dev.Touch = 0


    def swipe_direction(self, event):
        xdist = event["x1"] - event["x2"]
        ydist = event["y1"] - event["y2"]

        if abs(xdist) > abs(ydist):
            return "right" if xdist < 0 else "left"
        else:
            return "down" if ydist < 0 else "up"