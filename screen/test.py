from core.font import font
from core.screen import Screen
from screen.clock import ClockScreen


class TestScreen(Screen):
    roboto = font()
    arial = font("arial.ttf", 10)
    style = [0, 1]

    def start(self):
        self.draw_page()
        self.render(False)
        super().start()
    

    def accept_event(self, event):
        if event["swipe"]:
            if event["direction"] == "left" and event["x1"] > self.image.width - 10:
                self.app.busy = True
                self.app.change_screen("clock")
        else:
            if event["duration"] > 1:
                self.app.busy = True
                self.app.change_screen("clock")
                return
            
            self.app.busy = True
            if event["x"] < 125:
                self.style[0] = 0 if self.style[0] else 1
            else:
                self.style[1] = 0 if self.style[1] else 1
            self.draw_page()
            self.render()
            self.app.busy = False


    def draw_page(self):
        self.reset_image()
        self.draw.rectangle([0, 0, self.image.width, 125], fill = self.style[0])
        self.draw.text(((self.image.width / 2) - 2, self.image.height / 2), "Assalaamu", font = self.roboto, fill = 0 if self.style[0] else 1, anchor = "rm")
        self.draw.rectangle([125, 0, self.image.width, 125], fill = self.style[1])
        self.draw.text(((self.image.width / 2) + 2, self.image.height / 2), "'alaykum!", font = self.roboto, fill = 0 if self.style[1] else 1, anchor = "lm")
