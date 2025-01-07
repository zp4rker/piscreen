from PIL import Image, ImageDraw
from core.font import font
from datetime import datetime
from hijridate import Hijri, Gregorian

arial10 = font("arial.ttf", 10)

def top_bar():
    image = Image.new("1", (250, 13), 255)
    draw = ImageDraw.Draw(image)
    t = datetime.now()
    h = Gregorian.fromdate(t.date()).to_hijri()
    draw.rectangle((0, 0, image.width, 12), fill = 1)
    draw.text((2, 2), t.strftime("%d-%m-%Y"), font = arial10, fill = 0, anchor = "lt")
    draw.text((image.width / 2, 2), t.strftime("%H:%M"), font = arial10, anchor = "mt")
    draw.text((image.width - 2, 2), h.dmyformat("-"), font = arial10, anchor = "rt")
    draw.line([0, 12, image.width, 12], fill = 0)
    return image