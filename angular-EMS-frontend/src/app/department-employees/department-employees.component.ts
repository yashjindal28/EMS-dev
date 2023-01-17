import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { AuthTokenService } from '../auth-token.service';
import { Departments, Employee, ProjectEmployees,Projects } from '../employee';
import { EmployeeService } from '../employee.service';

@Component({
  selector: 'app-department-employees',
  templateUrl: './department-employees.component.html',
  styleUrls: ['./department-employees.component.css']
})
export class DepartmentEmployeesComponent implements OnInit {

  department : Departments = new Departments();
  projects: Projects = new Projects;
  employees: Employee[];
  eid: string | null

  constructor(private route: ActivatedRoute,
    public employeeService: EmployeeService, private router: Router,private authTokenService: AuthTokenService) { }

  ngOnInit(): void {
    this.eid = this.authTokenService.getEmployeeID()
    this.department.did = this.route.snapshot.params['did']
    this.employeeService.getDepartmentByID(this.eid , this.department.did).subscribe(data => {
      this.department = data;
    });

  

  }

  isAssigned(eid: string ): string{
    var result!: string
      this.department.eiddpt.forEach( (element) => {
        if( eid == element.eid){
          result =  element.assignedTo
        }
    }); 
  
  return result;
 
  }

  onSubmit(eid: string){
    var pid!: string
    var projectName = (document.getElementById('projectName') as HTMLInputElement).value
    this.department.projects.forEach( (element) => {
      if( projectName == element.projectName){
        pid = element.pid
      }
    });
    let customObj = { pid: pid , eid: eid , did: this.department.did}
    
    this.employeeService.saveEmployeeToAProject(this.eid,customObj).subscribe(data => {
        
    });
    
    this.router.routeReuseStrategy.shouldReuseRoute = () => false;
    this.router.onSameUrlNavigation = 'reload';
    this.router.navigate(['department-employees',this.department.did])
  }

}
