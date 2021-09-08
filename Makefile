env:
	@python -m venv env

requirements:
	@pip install -r requirements.txt

run:
	@uvicorn app.server:server --reload