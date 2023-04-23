const componentOne = () => {
    let e = document.createElement('div');
    e.setAttribute('id', 'componentOne');
    let e0 = document.createElement('div');
    let e1 = document.createElement('h1');
    e1.textContent = 'Some text';
    e0.appendChild(e1);
    e.appendChild(e0);
    return e;
};
const componentThree = () => {
    let e = document.createElement('div');
    e.setAttribute('id', 'componentThree');
    let e2 = componentOne() let e3 = document.createElement('h2');
    e3.textContent = 'Some Text';
    e2.appendChild(e3);
    e.appendChild(e2);
    return e;
};
const componentTwo = () => {
    let e = document.createElement('div');
    e.setAttribute('id', 'componentTwo');
    let e4 = document.createElement('div');
    let e5 = document.createElement('button');
    e5.textContent = 'Click me';
    e4.appendChild(e5);
    e.appendChild(e4);
    return e;
};
