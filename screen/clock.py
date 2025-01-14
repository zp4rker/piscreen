from datetime import datetime, timedelta
from threading import Thread

from core.font import font
from core.screen import Screen
from screen import components


class ClockScreen(Screen):
    style = 0
    roboto = font(size = 48)
    arial = font("arial.ttf", 10)
        

    def start(self):
        self.full_screen = True
        self.draw_page()
        self.render(False)
        loop = Thread(target=self.time_loop)
        loop.start()
        super().start()


    def draw_page(self):
        self.reset_image()
        t = datetime.now()
        self.draw.rectangle([0, 0, self.image.width, self.image.height], fill = 0 if self.style else 1)
        self.draw.text((self.image.width/2, self.image.height/2), t.strftime("%H:%M"), font = self.roboto, fill = self.style, anchor = "mm")


    def accept_event(self, event):
        if event["swipe"]:
            if event["direction"] == "right" and event["x1"] < 10:
                self.app.busy = True
                self.app.change_screen("test")
        else:
            self.app.busy = True
            self.style = 0 if self.style else 1
            self.draw_page()
            self.render()
            self.app.busy = False


    def time_loop(self):
        while self.app.current_screen == self and self.app.running:
            self.draw_page()
            self.render()
            t = datetime.now().replace(microsecond = 0, second = 0) + timedelta(minutes = 1)
            while t > datetime.now() and self.app.running:
                pass