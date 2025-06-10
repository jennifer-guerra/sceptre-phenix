#!/bin/bash
docker build -t phenix -f docker/Dockerfile --build-arg INSTALL_CERTS=https://raw.github.nrel.gov/Operations-Support/Trust-NREL-CA/master/NREL-CA/nrel-ca.pem .
cd docker/jit/
docker build -t phenix-jit -f Dockerfile --build-arg INSTALL_CERTS=https://raw.github.nrel.gov/Operations-Support/Trust-NREL-CA/master/NREL-CA/nrel-ca.pem .
