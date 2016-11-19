captcha.so
==========

File Size: about 3.5 MB

Replace captcha generator in your (Rails) project with a small .so file built with a Golang package
[github.com/dchest/captcha](https://github.com/dchest/captcha).

No need to install any image library.

This function generates a base64 encoded string of the captcha png file, which can be used as the data URI in HTML.

Example
=======

```ruby
require 'fiddle'

module Captcha
  def self.new(code, width = 240, height = 80)
    lib_captcha = Fiddle.dlopen('./captcha.so')
    new_captcha = Fiddle::Function.new(
      lib_captcha['NewCaptcha'],
      [Fiddle::TYPE_VOIDP, Fiddle::TYPE_VOIDP, Fiddle::TYPE_INT, Fiddle::TYPE_INT],
      Fiddle::TYPE_VOIDP,
    )
    img = new_captcha.call(SecureRandom.hex(10), code.to_s, width, height).to_s
    img.start_with?('iV') ? img : nil
  end
end


# in your controller:

number = Array.new(4).map{ SecureRandom.random_number(10) }.join
image = Captcha.new(number, 200, 80)
session[:captcha] = number
render json: { image: image }

# javascript:
# $('#captcha-image').attr('src', 'data:image/png;base64,' + image);
```
