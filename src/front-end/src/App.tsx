import './App.css';
import {
  BrowserRouter as Router,
  Route,
  Link,
  Routes
} from 'react-router-dom';
import {
  Navbar,
  Nav,
  Container,
  Image,
} from 'react-bootstrap';
import { Sched } from './components/Sched';
import { FindRoom } from './components/API';

function App() {
  return (
    <Router>
      <Navbar bg="dark" expand="lg" variant="dark">
        <Container fluid>
          <Navbar.Brand href="/">
            <Image fluid src={require("./images/empty-room-logo.png")} width="150px" alt="empty-room logo" />
          </Navbar.Brand>
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
        <Route path="/" element={<FindRoom />} />
      </Routes>

    </Router>
  );
}

export const Map = () => {
  return (
    <Container className="App">
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



export default App;