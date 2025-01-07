from core.screen import Screen
from screen.clock import ClockScreen
from core.font import font

class TestScreen(Screen):
    font = font()
    style = [0, 1]

    def start(self):
        self.draw_page()
        self.app.draw(self.image, partial = False)
    

    def accept_event(self, event):
        if event["swipe"]:
            return
        
        if event["duration"] > 1:
            self.app.busy = True
            self.app.change_screen(ClockScreen(self.app))
            return
        
        self.app.busy = True
        if event["x"] < 125:
            self.style[0] = 0 if self.style[0] else 1
        else:
            self.style[1] = 0 if self.style[1] else 1
        self.draw_page()
        self.app.draw(self.image)
        self.app.busy = False


    def draw_page(self):
        self.reset_image()
        self.draw.rectangle([0, 0, self.image.width, 125], fill = self.style[0])
        self.draw.text((125 - 2, 61), "Assalaamu", font = self.font, fill = 0 if self.style[0] else 1, anchor = "rm")
        self.draw.rectangle([125, 0, self.image.width, 125], fill = self.style[1])
        self.draw.text((125 + 2, 61), "'alaykum!", font = self.font, fill = 0 if self.style[1] else 1, anchor = "lm")
