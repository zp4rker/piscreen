from core.screen import Screen
from datetime import datetime
from threading import Thread
from core.font import font

class ClockScreen(Screen):
    style = 0
    font = font(48)

    def start(self):
        self.app.draw(self.image, partial = False)
        loop = Thread(target=self.time_loop)
        loop.start()
        super().start()


    def draw_page(self):
        self.reset_image()
        t = datetime.now()
        self.draw.rectangle([0, 0, self.image.width, self.image.height], fill = 0 if self.style else 1)
        self.draw.text((self.image.width/2, self.image.height/2), t.strftime("%H:%M"), font = self.font, fill = self.style, anchor = "mm")


    def accept_event(self, event):
        if event["swipe"]:
            if event["direction"] == "right" and event["x1"] < 10:
                self.app.busy = True
                self.app.back_screen()
        else:
            self.app.busy = True
            self.style = 0 if self.style else 1
            self.draw_page()
            self.app.draw(self.image)
            self.app.busy = False


    def time_loop(self):
        while self.app.current_screen == self and self.app.running:
            self.draw_page()
            self.app.draw(self.image)
            t = datetime.now().replace(microsecond = 0, second = 0)
            t = t.replace(minute = t.minute + 1)
            while t > datetime.now() and self.app.running:
                pass