from PIL import Image, ImageDraw
from core.font import font
from datetime import datetime

arial10 = font("arial.ttf", 10)

def top_bar():
    image = Image.new("1", (250, 13), 255)
    draw = ImageDraw.Draw(image)
    t = datetime.now()
    draw.rectangle((0, 0, image.width, 12), fill = 1)
    draw.text((2, 2), t.strftime("%H:%M"), font = arial10, fill = 0, anchor = "lt")
    draw.line([0, 12, image.width, 12], fill = 0)
    return image