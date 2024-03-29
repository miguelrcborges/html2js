/**
 * Generates componentOne component.
 * @return Component
 */
const componentOne=()=>{const e=document.createElement('div');e.setAttribute('id','componentOne');const e0=document.createElement('div');e0.setAttribute('class','epic');const e1=document.createElement('h1');e1.textContent=`Some text`;e0.appendChild(e1);e.appendChild(e0);return e;};
/**
 * Generates componentThree component.
 * @return Component
 */
const componentThree=()=>{const e=document.createElement('div');e.setAttribute('id','componentThree');const e0=componentOne();const e1=document.createElement('h2');e1.textContent=`Some Text`;e0.appendChild(e1);e.appendChild(e0);return e;};
/**
 * Generates componentFour component.
 * @param {string} name The dood's name
 * @return Component
 */
const componentFour=(name)=>{const e=document.createElement('div');e.setAttribute('id','componentFour');const e0=document.createElement('h1');e0.textContent=`Some random test with "'\\\`.`;e.appendChild(e0);const e1=document.createElement('p');e1.textContent=`Hello ${name}`;e.appendChild(e1);return e;};
/**
 * Generates componentTwo component.
 * @return Component
 */
const componentTwo=()=>{const e=document.createElement('div');e.setAttribute('id','componentTwo');const e0=document.createElement('div');e0.setAttribute('class','other-class');const e1=document.createElement('button');e1.setAttribute('id','button-2');e1.textContent=`Click me`;e0.appendChild(e1);e.appendChild(e0);return e;};
