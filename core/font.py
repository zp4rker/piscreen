from PIL import ImageFont

def font(file="roboto.ttf", size=14):
    return ImageFont.truetype(file, size)