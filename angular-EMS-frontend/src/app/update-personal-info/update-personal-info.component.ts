import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { PersonalInfo } from '../employee';
import { EmployeeService } from '../employee.service';

@Component({
  selector: 'app-update-personal-info',
  templateUrl: './update-personal-info.component.html',
  styleUrls: ['./update-personal-info.component.css']
})
export class UpdatePersonalInfoComponent implements OnInit {

  personalInfo: PersonalInfo = new PersonalInfo();
  eid: string
  constructor(private employeeService: EmployeeService,
    private route: ActivatedRoute,private router:Router) { }

  ngOnInit(): void {
    this.eid = this.route.snapshot.params['eid']
    this.employeeService.getPersonalInfoById(this.eid).subscribe( data => {
      this.personalInfo = data;
    });

  }

  onSubmit(){
    this.employeeService.updatePersonalInfoById(this.eid,this.personalInfo).subscribe( data => {
      this.goToViewPersonalInfo()
    }, 
    error => console.log(error));
  }

  goToViewPersonalInfo(){
    this.router.navigate(['/personal',this.eid]);
  }
}
