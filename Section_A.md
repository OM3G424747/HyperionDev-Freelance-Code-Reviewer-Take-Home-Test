<H2>Correctness</H2>

<p>It seems there are several error that's preventing your code from running.</p>
<p>I'm going to cover each of them and suggest areas of improvement, but it would be a good idea to always test your code before submitting it and make sure the most recent version is submitted.</p>

<H3>Indentation</H3>
It seems your indentation is incorrect from lines 2 to 6, and then lines 8 and 10.
When running python code it's always important to ensure identation is correct.
The easiest way to do this, is by pressing "TAB" instead of pressing the spacebar multiple times.
If you are coding on a text editor like Nano on Linux, please make sure you configure it to add 4 spaces instead if 2.
If you're using Visual Studio Code, I would recommend using an extension like "<strong><a href="https://marketplace.visualstudio.com/items?itemName=oderwat.indent-rainbow">indent-rainbow</a></strong>"


<H2>Efficiency</H2>
<H2>Style</H2>
<H2>Documentation</H2>

# Foobar

Foobar is a Python library for dealing with word pluralization.

## Installation

Use the package manager [pip](https://pip.pypa.io/en/stable/) to install foobar.

```bash
pip install foobar
```

## Usage

```python
import foobar

# returns 'words'
foobar.pluralize('word')

# returns 'geese'
foobar.pluralize('goose')

# returns 'phenomenon'
foobar.singularize('phenomena')
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)