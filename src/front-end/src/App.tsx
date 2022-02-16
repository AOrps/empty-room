import Component from 'react';
import './App.css';
import {
  Inject, ScheduleComponent,
  Day, Week, WorkWeek, Month, Agenda,
  EventSettingsModel,
  ViewsDirective,
  ViewDirective
} from '@syncfusion/ej2-react-schedule';
import { BrowserRouter as Router, Route, Link, Redirect, Switch } from 'react-router-dom';
import { Navbar, Nav } from 'react-bootstrap';




function App() {
  return (
    <Router>
      <Navbar bg="light" expand="lg">
        <Navbar.Brand href="/"><span className="sigmal-color">empty-room</span></Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
          <Nav className="me-auto">
            <Nav.Link as={Link} to={"/"}>Home</Nav.Link>
            {/* <Nav.Link as={Link} to={"/schedule"}>Schedule</Nav.Link>
            <Nav.Link href="http://github.com/AOrps/SigMal">Repository</Nav.Link>
            <Nav.Link as={Link} to={"/res"}>Resources</Nav.Link>  */}
          </Nav>
        </Navbar.Collapse>

      </Navbar>
      <div className="App">
        <header className="App-header">
          <p>
            Edit <code>src/App.tsx</code> and save to reload.
          </p>
          <div >
            <ScheduleComponent
              currentView='Week'
              readonly={true}
            >
              <ViewsDirective>
                <ViewDirective
                  option='Week'
                  startHour='6:00'
                  endHour='24:00'
                />
              </ViewsDirective>

              <Inject services={[Day, Week, WorkWeek, Month, Agenda]} />
            </ScheduleComponent>
          </div>
        </header>
      </div>

      <Switch>
        {/* <Route path="/schedule">
              <h1 className="objpad-top">Schedule</h1>
              <ScheduleMonthView />
            </Route> */}
        <Route path="/map">
          <Redirect
            to={{
              pathname: "https://www.njit.edu/sites/default/files/NJIT-CAMPUS-MAP-CURRENT-WEST-VIEW.jpg"
            }}
          />
        </Route>
        <Route path="/">
          <App />
        </Route>
      </Switch>


    </Router>
  );
}

export default App;