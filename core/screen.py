from PIL import ImageDraw, Image

class Screen(object):
    def __init__(self, app):
        self.app = app
        self.reset_image()

    
    def reset_image(self):
        self.image = Image.new("1", (self.app.epd.height, self.app.epd.width), 255)
        self.draw = ImageDraw.Draw(self.image)
        self.draw.rectangle([(0, 0), (self.app.epd.height, self.app.epd.width)], fill = 1)


    def start(self):
        self.app.busy = False


    def accept_event(self, event):
        pass