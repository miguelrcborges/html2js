const componentOne = () => {
    let e = document.createElement('div');
    e.setAttribute('id', 'componentOne');
    let e0 = document.createElement('div');
    e0.classList.add('epic');
    let e1 = document.createElement('h1');
    e1.textContent = 'Some text';
    e0.appendChild(e1);
    e.appendChild(e0);
    return e;
};
const componentThree = () => {
    let e = document.createElement('div');
    e.setAttribute('id', 'componentThree');
    let e0 = componentOne();
    let e1 = document.createElement('h2');
    e1.textContent = 'Some Text';
    e0.appendChild(e1);
    e.appendChild(e0);
    return e;
};
const componentTwo = () => {
    let e = document.createElement('div');
    e.setAttribute('id', 'componentTwo');
    let e0 = document.createElement('div');
    e0.classList.add('other-class');
    let e1 = document.createElement('button');
    e1.setAttribute('id', 'button-2');
    e1.textContent = 'Click me';
    e0.appendChild(e1);
    e.appendChild(e0);
    return e;
};
