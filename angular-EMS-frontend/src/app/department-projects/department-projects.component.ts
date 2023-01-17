import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { AuthTokenService } from '../auth-token.service';
import { Departments, Projects } from '../employee';
import { EmployeeService } from '../employee.service';

@Component({
  selector: 'app-department-projects',
  templateUrl: './department-projects.component.html',
  styleUrls: ['./department-projects.component.css']
})
export class DepartmentProjectsComponent implements OnInit {

  
  department: Departments = new Departments();
  projectName: string;
  pid: string;
  deadline: string;
  projectList: Array<Projects> = [];
  
  eid: string | null //eid of manager who has logged in to pass for auth verification 

  constructor(private employeeService: EmployeeService,
    private route: ActivatedRoute,private router: Router,private authTokenService: AuthTokenService) { }

  ngOnInit(): void {
    this.eid = this.authTokenService.getEmployeeID()
    this.department.did = this.route.snapshot.params['did']
    this.employeeService.getDepartmentByID(this.eid , this.department.did).subscribe(data => {
      this.department = data;
    
    });


  }

  onSubmit(){
    let customObj = new Projects();
    customObj.projectName = (document.getElementById('projectName') as HTMLInputElement).value
    customObj.pid = (document.getElementById('pid') as HTMLInputElement).value
    customObj.deadline = (document.getElementById('deadline') as HTMLInputElement).value
    console.log(customObj)
    this.employeeService.saveProjectInDepartmentOfID(this.eid,this.department.did,customObj).subscribe(data => {
        
    });
    this.router.routeReuseStrategy.shouldReuseRoute = () => false;
    this.router.onSameUrlNavigation = 'reload';
    this.router.navigate(['department-projects',this.department.did])
  }

}
