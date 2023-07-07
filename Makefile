spanner:
	podman pull gcr.io/cloud-spanner-emulator/emulator:1.3.0	 

start-spanner:
	podman run -p 9010:9010 --name spanner-emulator -d gcr.io/cloud-spanner-emulator/emulator:1.3.0

compose-up:
	podman-compose up

compose-down:
	podman-compose down

terminal:
	podman-compose exec spanner-cli spanner-cli -p spanner-project -i spanner-instance -d spanner-database


.PHONY:spanner start-spanner compose-up terminal