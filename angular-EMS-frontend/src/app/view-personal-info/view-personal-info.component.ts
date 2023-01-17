import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Employee, PersonalInfo } from '../employee';
import { EmployeeService } from '../employee.service';

@Component({
  selector: 'app-view-personal-info',
  templateUrl: './view-personal-info.component.html',
  styleUrls: ['./view-personal-info.component.css']
})
export class ViewPersonalInfoComponent implements OnInit {

  employee : Employee = new Employee();
  personalInfo : PersonalInfo = new PersonalInfo();
  eid : string;
  constructor(private employeeService : EmployeeService,private router: Router,
    private route: ActivatedRoute) { }

  ngOnInit(): void {
    this.eid = this.route.snapshot.params['eid']

    this.employeeService.getEmployeeById(this.eid,this.eid).subscribe( data => {
      this.employee = data;
    });

    this.employeeService.getPersonalInfoById(this.eid).subscribe( data => {
      this.personalInfo = data;
    });

  }

  updatePersonalInfo(eid : string){
    this.router.navigate(['update-personal',eid]);
  }

}
