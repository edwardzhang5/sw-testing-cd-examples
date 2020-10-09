// import express from 'express';
// import { urlencoded, json } from 'body-parser';
// import titleCase from './titlecase';

const express = require('express');
const { urlencoded, json } = require('body-parser');
const titleCase = require('./titlecase');

const app = express();

app.use(urlencoded({ extended: false }));
app.use(json());

app.get('/', (req, res) => {
  res.status(200).send('Movie API');
});

app.post('/movietitle', (req, res) => {
  const title = titleCase(req.body.title);
  res.status(200).send({ title });
});

app.get('/movietitle/:title', (req, res) => {
  const title = titleCase(req.params.title);
  res.status(200).send({ title });
});

const _app = app;
module.exports.app = _app;
