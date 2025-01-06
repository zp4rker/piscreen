from PIL import ImageDraw, Image

class Screen(object):
    def __init__(self, app):
        self.reset_image(app)
        self.draw.rectangle([(0,0),(app.epd.height, app.epd.width)], fill = 1)

    
    def reset_image(self, app):
        self.image = Image.new("1", (app.epd.height, app.epd.width), 255)
        self.draw = ImageDraw.Draw(self.image)


    def start(self, app):
        pass


    def accept_event(self, app, event):
        pass