from core.screen import Screen
from screen.clock import ClockScreen

class TestScreen(Screen):
    style = [0, 1]

    def start(self, app):
        self.draw_page(app)
        app.draw(self.image, partial = False)
    

    def accept_event(self, app, event):
        if event["swipe"]:
            return
        
        if event["duration"] > 1:
            app.change_screen(ClockScreen(app))
            return
        
        app.busy = True
        if event["x"] < 125:
            self.style[0] = 0 if self.style[0] else 1
        else:
            self.style[1] = 0 if self.style[1] else 1
        self.draw_page(app)
        app.draw(self.image)
        app.busy = False


    def draw_page(self, app):
        self.reset_image(app)
        self.draw.rectangle([0, 0, app.epd.height, 125], fill = self.style[0])
        self.draw.text((125 - 2, 61), "Assalaamu", font = app.font, fill = 0 if self.style[0] else 1, anchor = "rm")
        self.draw.rectangle([125, 0, app.epd.height, 125], fill = self.style[1])
        self.draw.text((125 + 2, 61), "'alaykum!", font = app.font, fill = 0 if self.style[1] else 1, anchor = "lm")
