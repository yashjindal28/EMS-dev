import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ManagerUpdateEmployeeComponent } from './manager-update-employee.component';

describe('ManagerUpdateEmployeeComponent', () => {
  let component: ManagerUpdateEmployeeComponent;
  let fixture: ComponentFixture<ManagerUpdateEmployeeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ManagerUpdateEmployeeComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ManagerUpdateEmployeeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
