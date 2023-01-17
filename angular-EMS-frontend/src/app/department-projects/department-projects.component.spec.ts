import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DepartmentProjectsComponent } from './department-projects.component';

describe('DeaprtmentProjectsComponent', () => {
  let component: DepartmentProjectsComponent;
  let fixture: ComponentFixture<DepartmentProjectsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DepartmentProjectsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(DepartmentProjectsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
