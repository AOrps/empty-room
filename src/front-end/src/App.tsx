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
  Form,
  Row, Col
} from 'react-bootstrap';
import { Sched } from './components/Sched';
import { InfoTable } from './components/Building';


function App() {
  return (
    <Router>
      <Navbar bg="dark" expand="lg" variant="dark">
        <Container fluid>
          <Navbar.Brand href="/">
            <Image fluid src={require("./images/empty-room-logo.png")} width="150px" alt="empty-room logo"/>
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

export const FindRoom = () => {

  return (
    <Container className="App">
      <Row>
        <Col>
          <br />
          <InfoTable />
        </Col>
        <Col>
        <br />
          <Form>
            <Form.Label>Enter Building to find an empty-room</Form.Label>
            <Form.Control id="building" />
            <Form.Text>It is totally valid to enter the Building Abbreviation</Form.Text>
          </Form>
        </Col>
      </Row>

    </Container>
  );
}


export default App;