import React from "react";
import 'bootstrap/dist/css/bootstrap.css';
//import logo from './logo.svg';
import './App.css';
import { Container } from 'semantic-ui-react';
import Tasks from './components/task';
function App() {
  return (
    <div>
         <Container>
          <Tasks/>
         </Container>
    </div>
    
  );
}

export default App;