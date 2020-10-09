const request = require('supertest');
const app = require('./app').app;

describe('server tests', () => {
  it('serves the root endpoint /', done => {
    request(app).get('/').expect(200).expect('Movie API').end(done);
  });

  it('serves the root with jest test', done => {
    request(app)
      .get('/')
      .then(res => {
        expect(res.statusCode).toBe(200);
        expect(res.text).toBe('Movie API');
        done();
      });
  });

  it('provied titlecase title for post request', done => {
    request(app)
      .post('/movietitle')
      .type('form')
      .send({ title: 'movie title' })
      .then(res => {
        console.log(res.body);
        expect(res.body).toHaveProperty('title', 'Movie Title');
        done();
      });
  });
});
