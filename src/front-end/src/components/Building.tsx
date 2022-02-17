import {
    Table
} from 'react-bootstrap';


export const InfoTable = () => {
    const columns: Array<String> = [ "Building By Name", "Building #", "Abbreviation"
        // { dataField: "name", text: "Name" },
        // { dataField: "number", text: "Building Number" },
        // { dataField: "abbr", text: "Abbreviations" },
    ];
    
    const building_data = [
        { building: "Campus Center", number: "12", abbreviation: "CTR | CC" },
        { building: "Central Avenue Building", number: "6", abbreviation: "CAB" },
        { building: "Central King Building", number: "7", abbreviation: "CKB" },
        { building: "Colton Hall", number: "5", abbreviation: "COLT" },
        { building: "Cullimore Hall", number: "10", abbreviation: "CULM" },
        // { building: "Eberhardt Hall", number: "9", abbreviation: "EBER" },
        { building: "Electrical and Computer Eng. Center", number: "16", abbreviation: "ECEC" },
        { building: "Faculty Memorial Hall", number: "15", abbreviation: "FMH" },
        { building: "Fenster Hall", number: "8", abbreviation: "FENS" },
        { building: "Guttenberg Information Technology Center", number: "31", abbreviation: "GITC" },
        { building: "Kupfrian Hall", number: "13", abbreviation: "KUPF" },
        // { building: "Life Sciences & Engineering Center", number: "2a", abbreviation: "LSEC" },
        { building: "Mechanical Engineering Center", number: "30", abbreviation: "ME" },
        // { building: "Microelectronics Center", number: "14", abbreviation: "MIC" },
        // { building: "Specht Building", number: "3", abbreviation: "SPEC" },
        { building: "Tiernan Hall", number: "26", abbreviation: "TIER" },
        { building: "Weston Hall", number: "4", abbreviation: "WEST" },
        // { building: "York Center", number: "2", abbreviation: "YORK" },
    ];
    
    return (
        <Table striped bordered hover size="sm"> {/*variant="dark"> */}
        <thead>
          <tr>
            <th>Building By Name</th>
            <th>Building #</th>
            <th>Abbreviation</th>
          </tr>
        </thead>
        <tbody>
        {Array.from(building_data).map((val) =>( 
          <tr>
                <td> {val.building} </td>
                <td> {val.number} </td>
                <td> {val.abbreviation} </td>        
          </tr>
        ))}
        </tbody>
      </Table>
    );

};