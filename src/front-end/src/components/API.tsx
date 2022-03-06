import {
    Container,
    Form,
    Button,
    Row, Col, FormControl
} from 'react-bootstrap';
import { InfoTable } from './Building';
import React, { Component } from "react";
import axios, { Axios } from "axios";

class FindRoom extends Component<{}, any> {

    constructor(props:any) {
        super(props);
        this.state = {
            building: "",
            day: "",
            // result: ""
        };
        this.onInputchange = this.onInputchange.bind(this);
        this.onSubmitForm = this.onSubmitForm.bind(this);
    }

    onInputchange = (event) => {
        this.setState({
          [event.target.name]: event.target.value
        });
      }

      onSubmitForm = (event) => {
            event.preventDefault();
            // console.log(this.state);
            this.handleButtonClick();
      }

    handleButtonClick = () => {
        const fetchUserEmail = async () => {
            const response = axios.post("https://ec2-18-119-118-48.us-east-2.compute.amazonaws.com",
                {
                    building: this.state.building,
                    day: this.state.day
                }
                // method: 'POST',
                // JSON.stringify({
                    // building: this.state.building,
                    // day: this.state.day,
                // })
            ).then(response => {
                console.log(response)
                this.setState({result: response})
                // alert(response.json())
            });
            // const { result } = await response.json();
            // this.setState({
            //     result
            // });
        };
        console.log(this.state);
        fetchUserEmail();
    };


    render() {
        const { building, day, result } = this.state;

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
                        <Form>
                            <Row>
                                <Col>
                                    <Form.Label>Enter Building</Form.Label>
                                    <FormControl id="building" name="building" value={this.state.building} onChange={this.onInputchange} />
                                    <Form.Text>You can use building Abbreviations</Form.Text>
                                </Col>
                                <Col>
                                    <Form.Label>Enter Day</Form.Label>
                                    <FormControl id="day" name="day" value={this.state.day} onChange={this.onInputchange} />
                                    <Form.Text>ex. M -&gt; Monday, T -&gt; Tuesday, W -&gt; Wednesday</Form.Text>
                                </Col>
                            </Row>
                            <br />
                            <Row>
                                <Button variant="outline-info" type='submit' onClick={this.onSubmitForm}>Find empty-room!</Button>
                            </Row>
                            <br />
                            <div>{this.state.result}</div>
                        </Form>
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