import { Component, OnInit } from '@angular/core';
import { FormGroup, NgForm } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EmployeeLeave } from '../employee';
import { EmployeeService } from '../employee.service';

@Component({
  selector: 'app-leave-employee',
  templateUrl: './leave-employee.component.html',
  styleUrls: ['./leave-employee.component.css']
})
export class LeaveEmployeeComponent implements OnInit {

  employeeLeaveData : EmployeeLeave = new EmployeeLeave()
  eid : string
  
  constructor(private employeeService: EmployeeService,
    private route: ActivatedRoute,private router:Router) { }

  ngOnInit(): void {
    this.eid = this.route.snapshot.params['eid']
    this.employeeService.getEmployeeLeaveDataByID(this.eid).subscribe(data => {
      this.employeeLeaveData = data;
    
    });
  }

  onSubmit(){
    this.employeeLeaveData.sdate = (document.getElementById('sdate') as HTMLInputElement).value
    this.employeeLeaveData.edate = (document.getElementById('edate') as HTMLInputElement).value
    this.employeeLeaveData.daysReq = parseFloat((document.getElementById('daysReq') as HTMLInputElement).value)
    this.employeeLeaveData.reason = (document.getElementById('reason') as HTMLInputElement).value
    
    this.employeeService.requestLeave(this.eid,this.employeeLeaveData).subscribe(data => {
      
    });
    this.router.routeReuseStrategy.shouldReuseRoute = () => false;
    this.router.onSameUrlNavigation = 'reload';
    this.router.navigate(['leave-employee',this.eid])
  }

}
