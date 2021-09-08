# SecTool

Application built for the third and last integration project of Computer Engineering.

## How to Run

The docker option is prefereed.

### Linux:

```sh
  # Creates a virtual environment
  make env

  # Activates the virtual environment
  source env/bin/activate

  # Install project requirements
  make requirements

  # Runs the application
  make run
```

### Windows:

```ps1
  # Creates a virtual environment
  py -m venv env

  # Activates the virtual environment
  .\env\Scripts\activate

  # Install project requirements
  pip install -r requirements.txt

  # Runs the application
  uvicorn app.server:server --reload
```

### Docker:

Follow the specific OS instructions to install the required packages on a virtualenv. **DO NOT** run the "Runs the application" stage.

```sh
  # Runs the application
  docker-compose up -d
```

## Debugging

VSCode's debug tool is great, and I already left the debug script ready, so just do <kbd>Ctrl</kbd> <kbd>+</kbd> <kbd>Shift</kbd> <kbd>+</kbd> <kbd>P</kbd> and type `Start debugging`. Breakpoints and everything will work as expected.
