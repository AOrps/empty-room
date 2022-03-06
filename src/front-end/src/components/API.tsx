import {
    Container,
    Form,
    Button,
    Row, Col, FormControl
} from 'react-bootstrap';
import { InfoTable } from './Building';
import React, { Component } from "react";

class FindRoom extends Component {

    constructor() {
        super();
        this.state = {value: ""};
        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleChange(event: React.ChangeEvent) { this.setState({ value: event.target.value }); }
    handleSubmit(event) {
        alert('A name was submitted: ' + this.state.value);
        event.preventDefault();
    }


    render() {


        return (
            <Container className="App" >
                <Row>
                    <Col>
                        <br />
                        <InfoTable />
                    </Col>
                    <Col>
                        <br />
                        <h2>To Find an empty-room, fill out the following: </h2>
                        <form onSubmit={this.handleSubmit}>
                            <label>
                                Name:
                                <input type="text" value={this.state.value} onChange={this.handleChange} />        </label>
                            <input type="submit" value="Submit" />
                        </form>
                        <div>{this.state.value}</div>
                        {/* <Form>
                            <Row>
                                <Col>
                                    <Form.Label>Enter Building</Form.Label>
                                    <Form.Control id="building" />
                                    <Form.Text>You can use building Abbreviations</Form.Text>
                                </Col>
                                <Col>
                                    <Form.Label>Enter Day</Form.Label>
                                    <FormControl id="day" />
                                    <Form.Text>ex. M -&gt; Monday, T -&gt; Tuesday, W -&gt; Wednesday</Form.Text>
                                </Col>
                            </Row>
                            <br />
                            <Row>
                                <Button variant="outline-info" type='submit'>Find empty-room!</Button>
                            </Row>
                            <br />
                            <div></div>
                        </Form> */}
                    </Col>
                </Row>

            </Container>
        )
    };
}

// export const FindRoom = () => {

//     const [building,setBuilding] = React.useState("");
//     const [day,setDay] = React.useState("");
//     const [result, setResult] = React.useState("");

//     const requestOptions = {
//       method: 'POST',
//       // headers: { 'Content-Type': 'application/json' },
//       body: JSON.stringify({ 
//         buidling: {building},
//         day: {day},
//       })
//     };

//     const fetchData = async () => {
//       const response = await fetch("ec2-18-119-118-48.us-east-2.compute.amazonaws.com", requestOptions);
//       const { result } = await response.json();
//       // console.log(result);
//       setResult(result);
//     }


//     return (
//       <Container className="App">
//         <Row>
//           <Col>
//             <br />
//             <InfoTable />
//           </Col>
//           <Col>
//             <br />
//             <h2>To Find an empty-room, fill out the following: </h2>
//             <Form>
//               <Row>
//                 <Col>
//                   <Form.Label>Enter Building</Form.Label>
//                   <Form.Control id="building" />
//                   <Form.Text>You can use building Abbreviations</Form.Text>
//                 </Col>
//                 <Col>
//                   <Form.Label>Enter Day</Form.Label>
//                   <FormControl id="day" />
//                   <Form.Text>ex. M -&gt; Monday, T -&gt; Tuesday, W -&gt; Wednesday</Form.Text>
//                 </Col>
//               </Row>
//               <br />
//               <Row>
//                 <Button variant="outline-info" type='submit' onClick={() => {
//                   setBuilding(building);
//                   setDay(day);

//                 }}>Find empty-room!</Button>
//               </Row>
//               <br />
//               <div>{day}</div>
//             </Form>
//           </Col>
//         </Row>

//       </Container>
//     );
//   }


// export const POSTSend = () => {
//     const requestOptions = {
//         method: 'POST',
//         headers: { 'Content-Type': 'application/json' },
//         body: JSON.stringify({ title: 'React POST Request Example' })
//     };
//     return fetch('https://jsonplaceholder.typicode.com/posts', requestOptions);
// }

export default FindRoom;