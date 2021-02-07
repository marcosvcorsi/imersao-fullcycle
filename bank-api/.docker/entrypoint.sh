#!/bin/bash

yarn
yarn typeorm migration:run
yarn console fixtures
yarn start:dev