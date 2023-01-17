
export class Employee {
    _id : string; 
    eid : string;
    firstname : string;
    lastname: string;
    gender: string;
    desg: string;
    did:string;
    dpt: string;
    manager:string;
    managerID:string;
    salary:string;
    email:string;
    location:string;
    isCheckIssued: number;

}

export class PersonalInfo{
    eid: string;
    address: string;
    contact: string ; 
}

export class EmployeeLeave { 
    eid : string;
    firstname : string;
    lastname: string;
    manager:string;
    managerID:string;
    leavesRem: number;
    sdate:string;
    edate:string;
    daysReq: number;
    reason: string;
    pendingStatus: number;
    approved: number;

}

export class ProjectEmployees{
    eid : string;
}

export class Projects{
    pid:string;
    projectName:string;
    deadline:string;
    projectEmployees: ProjectEmployees[];
    
}

export class EmployeesOfDpt{
    eid:string;
    assignedTo:string;
    firstname: string;
    lastname: string;
    desg: string;
    email:string;
}

export class Departments {
    did: string;
    dpt: string;
    manager: string;
    managerID: string;
    eiddpt: EmployeesOfDpt[];
    projects: Projects[];
}

export class User{
    eid:string;
    password: string;
}

export class Count{
    countOfEmp:number;
    countOfDpt: number;
}

