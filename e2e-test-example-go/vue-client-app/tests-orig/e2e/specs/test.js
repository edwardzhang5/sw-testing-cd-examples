// https://docs.cypress.io/api/introduction/api.html

describe('SW Testing App Tests', () => {
  it('shows the body mass index page', () => {
    cy.visit(Cypress.env('BASE_URL') + '/bmi');
    cy.get('#bmi-title').contains('Body Mass Index');
  });

  it('Splits the tips among the number of guests correctly', () => {
    cy.visit(Cypress.env('BASE_URL') + '/');
    cy.get('input#cost').type('90.00');
    cy.get('input#guests').type('3');
    cy.get('#submit').click();
    cy.get('ul li:first').contains('$34.5');
  });
});
