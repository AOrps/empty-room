import {
    Table
} from 'react-bootstrap';

type BuildingInformation = {
    building: string;
    number: string;
    abbreviation: string;
}

export const InfoTable = () => {
    const columns: Array<String> = [ "Building By Name", "Building #", "Abbreviation"];

    const building_data: Array<BuildingInformation> = [
        { building: "Campus Center", number: "12", abbreviation: "CTR | CC" },
        { building: "Central Avenue Building", number: "6", abbreviation: "CAB" },
        { building: "Central King Building", number: "7", abbreviation: "CKB" },
        { building: "Colton Hall", number: "5", abbreviation: "COLT" },
        { building: "Cullimore Hall", number: "10", abbreviation: "CULM" },
        { building: "Electrical and Computer Eng. Center", number: "16", abbreviation: "ECEC" },
        { building: "Faculty Memorial Hall", number: "15", abbreviation: "FMH" },
        { building: "Fenster Hall", number: "8", abbreviation: "FENS" },
        { building: "Guttenberg Information Technology Center", number: "31", abbreviation: "GITC" },
        { building: "Kupfrian Hall", number: "13", abbreviation: "KUPF" },
        { building: "Mechanical Engineering Center", number: "30", abbreviation: "ME" },
        { building: "Tiernan Hall", number: "26", abbreviation: "TIER" },
        { building: "Weston Hall", number: "4", abbreviation: "WEST" },
    ];
    
    return (
        <Table striped bordered hover size="sm">
        <thead>
          <tr>
            {Array.from(columns).map((val) => (
                <th>{val}</th>
            ))}
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