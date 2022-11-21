#!/bin/bash

cp Dockerfile ../../dist
cd ../../dist

docker build --tag "plutono/debtest" .

rm Dockerfile

docker run -i -t plutono/debtest /bin/bash
