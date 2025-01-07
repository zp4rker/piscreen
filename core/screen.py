from PIL import ImageDraw, Image
from screen import components

class Screen(object):
    def __init__(self, app):
        self.app = app
        self.reset_image()

    
    def reset_image(self):
        self.image = Image.new("1", (self.app.epd.height, self.app.epd.width - 12), 255)
        self.draw = ImageDraw.Draw(self.image)


    def start(self):
        self.app.busy = False


    def render(self, partial=True):
        image = Image.new("1", (self.app.epd.height, self.app.epd.width), 255)
        image.paste(im = components.top_bar(), box = (0, 0))
        image.paste(self.image, (0, 13))
        self.app.draw(image, partial)


    def accept_event(self, event):
        pass