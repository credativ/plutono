#!/bin/bash

cp Dockerfile ../../dist
cd ../../dist

docker build --tag "plutono/rpmtest" .

rm Dockerfile

docker run -i -t plutono/rpmtest /bin/bash
