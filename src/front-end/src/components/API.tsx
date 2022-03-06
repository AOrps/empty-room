import {
    Container,
    Form,
    Button,
    Row, Col, FormControl
} from 'react-bootstrap';
import { InfoTable } from './Building';
import React from "react";

// export interface ReqData {
//     building: string;
//     day: string;
// }


// export const get10 : () => Promise<number> = async () => {
//     return parseInt( (Math.random() * 10).toFixed(0));
// };

export const FindRoom = () => {

    const [building,setBuilding] = React.useState("");
    const [day,setDay] = React.useState("");
    const [result, setResult] = React.useState("");
  
    const requestOptions = {
      method: 'POST',
      // headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ 
        buidling: {building},
        day: {day},
      })
    };
  
    const fetchData = async () => {
      const response = await fetch("ec2-18-119-118-48.us-east-2.compute.amazonaws.com", requestOptions);
      const { result } = await response.json();
      // console.log(result);
      setResult(result);
    }
  
  
    return (
      <Container className="App">
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
                <Button variant="outline-info" type='submit' onClick={() => {
                  setBuilding(building);
                  setDay(day);
  
                }}>Find empty-room!</Button>
              </Row>
              <br />
              <div>{day}</div>
            </Form>
          </Col>
        </Row>
  
      </Container>
    );
  }
  
  // export const Bonus = () => {
  //   return (
  //     <Container className="App">
  //       <ApiInfo />
  //     </Container>
  //   );
  // }

export const POSTSend = () => {
    const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ title: 'React POST Request Example' })
    };
    return fetch('https://jsonplaceholder.typicode.com/posts', requestOptions);
}



// export const ApiInfo = () => {
    


//     React.useEffect(() => {
//         const fetchRes = async () => {
//             const response = await POSTSend();
//         }
//     })
    
//     return (
//         <div>{res}</div>
//     );
// }

// export async class ApiInfo extends React.Component {
//     val = await get10();
//     render() {
//         return <div>{this.val}</div>
//     };
// }


// export const ApiInfo = async () => render(<div>{get10}</div>);


// export const ApiInfo = async () => {
//     let value = await get10();
//     // return render(
//     //     <div>{value}</div>
//     // );
//     return React.createElement("div", null, value);
// } 


// export const ApiInfo: () => Promise<string> = async (props: ReqData) => {

//     const requestOptions = {
//         method: 'POST',
//         headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
//         body: JSON.stringify({
//             buidling: props.building,
//             day: props.day
//         })
//     };

//     // const {data: response} = Axios.post('ec2-18-119-118-48.us-east-2.compute.amazonaws.com',requestOptions);
//     // this.setState({data: data})
//     // {/* <div>{response.data}</div> */}

//     await fetch('ec2-18-119-118-48.us-east-2.compute.amazonaws.com', requestOptions)
//         .then(response => response.json());

//     return (
//         <div>Bruh</div>
//     );

//     // return React.createElement(
//     //     () => <div>The element</div>,
//     //     // {key: }
//     // );
// }
