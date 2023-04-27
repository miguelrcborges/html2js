const componentOne = () => {
    const e = document.createElement('div');
    e.setAttribute('id', 'componentOne');
    const e0 = document.createElement('div');
    e0.setAtrribute('class', 'epic');
    const e1 = document.createElement('h1');
    e1.textContent = 'Some text';
    e0.appendChild(e1);
    e.appendChild(e0);
    return e;
};
const componentThree = () => {
    const e = document.createElement('div');
    e.setAttribute('id', 'componentThree');
    const e0 = componentOne();
    const e1 = document.createElement('h2');
    e1.textContent = 'Some Text';
    e0.appendChild(e1);
    e.appendChild(e0);
    return e;
};
const componentTwo = () => {
    const e = document.createElement('div');
    e.setAttribute('id', 'componentTwo');
    const e0 = document.createElement('div');
    e0.setAtrribute('class', 'other-class');
    const e1 = document.createElement('button');
    e1.setAtrribute('id', 'button-2');
    e1.textContent = 'Click me';
    e0.appendChild(e1);
    e.appendChild(e0);
    return e;
};
