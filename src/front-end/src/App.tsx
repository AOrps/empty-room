import './App.css';
import {
  BrowserRouter as Router,
  Route,
  Link,
  Routes
} from 'react-router-dom';
import { Navbar, Nav, Container, Image } from 'react-bootstrap';
import { Sched } from './components/Sched';


function App() {
  return (
    <Router>
      <Navbar bg="dark" expand="lg" variant="dark">
        <Container fluid>
          <Navbar.Brand href="/">empty-room</Navbar.Brand>
          <Navbar.Toggle aria-controls="basic-navbar-nav" />
          <Navbar.Collapse id="basic-navbar-nav">
            <Nav className="me-auto">
              <Nav.Link as={Link} to={"/map"}>Map</Nav.Link>
              <Nav.Link as={Link} to={"/schedule"}>Schedule</Nav.Link>
              <Nav.Link as={Link} to={"/find-room"}>Find Room</Nav.Link>
            </Nav>
          </Navbar.Collapse>
        </Container>
      </Navbar>

      <Routes>
        <Route path="/map" element={<Map />} />
        <Route path="/schedule" element={<CustomSchedule />} /> 
        <Route path="/find-room" element={<FindRoom />} />
        <Route path="/" element={<CustomSchedule />} />
      </Routes>

    </Router>
  );
}

export const Map = () => {
  return (
    <Container>
    <Image fluid src={require("./images/njit.jpg")} alt="MAP of NJIT" />
    </Container>
  );
}

export const CustomSchedule = () => {
  return (
  <Container className="App">
    <Sched />
  </Container>
  );
}

export const FindRoom = () => {
  return (
    <Container>

    </Container>
  );
}


export default App;