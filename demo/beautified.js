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
const componentTwo = () => {
    let e = document.createElement('div');
    e.setAttribute('id', 'componentTwo');
    let e2 = document.createElement('div');
    let e3 = document.createElement('button');
    e3.textContent = 'Click me';
    e2.appendChild(e3);
    e.appendChild(e2);
    return e;
};
