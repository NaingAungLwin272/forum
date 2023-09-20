import { Component } from '@angular/core';
import { Team } from 'src/app/interfaces/team';
import { TeamService } from 'src/app/repositories/team.service';
import { DepartmentService } from 'src/app/repositories/department.service';
import { NzMessageService } from "ng-zorro-antd/message"
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { Department } from 'src/app/interfaces/department';
import { NzModalService } from 'ng-zorro-antd/modal';
import { TeamCreateComponent } from 'src/app/components/team-create/team-create.component';
import { LoaderService } from 'src/app/services/loader.service';
import { ApiService } from 'src/app/services/api.service';

@Component({
  selector: 'app-teams-list',
  templateUrl: './teams-list.component.html',
  styleUrls: ['./teams-list.component.scss']
})
export class TeamsListComponent {
  teamData: Team[] = []
  departmentData: Department[] = [];
  departmentName!: string;
  teamEditForm!: FormGroup;
  clickedId!: string | '';
  totalTeams: number = 0;
  pageIndex!: number;
  pageSize!: number;
  teamCount: number = 0;

  constructor(
    private teamSvc: TeamService,
    private departmentSvc: DepartmentService,
    private messageSvc: NzMessageService,
    private modalSvc: NzModalService,
    private formBuilder: FormBuilder,
    private loaderSvc: LoaderService,
    private apiSvc: ApiService
  ) {
    this.teamEditForm = this.formBuilder.group({
      name: ['', [Validators.required]],
      department_id: ['', [Validators.required]]
    });
  }

  ngOnInit(): void {
    this.pageIndex = 1;
    this.pageSize = 10;
    this.getTeamCount()
    this.getTeamLists(this.pageIndex, this.pageSize)
    this.getDepartmentList()
  }

  getTeamCount(): void {
    this.teamSvc.getTeamCount().subscribe({
      next: (teamCount) => {
        this.teamCount = teamCount.count;
      }, error: (error) => {
        console.log(error, 'error.....')
      }
    })
  }

  getTeamLists(pageIndex?: number, pageSize?: number): void {
    this.loaderSvc.call()
    this.teamSvc.getTeamList(pageIndex, pageSize).subscribe({
      next: (teamListData: Team[]) => {
        this.teamData = teamListData;
        this.updateDepartmentNames();
        this.loaderSvc.dissAll();
      },
      error: (error) => {
        this.loaderSvc.dissAll();
        this.teamSvc.handleErrorType(error)
      }
    })
  }

  getDepartmentList(): void {
    this.departmentSvc.getDepartmentList().subscribe({
      next: (departmentListData: Department[]) => {
        this.departmentData = departmentListData;
        this.updateDepartmentNames();
      },
      error: (error) => {
        console.log(error, 'error.....')
      }
    });
  }

  cancel(): void {
    this.messageSvc.info('click cancel');
  }

  confirm(id: string): void {
    this.loaderSvc.call();
    this.teamSvc.deleteTeam(id).subscribe({
      next: () => {
        this.messageSvc.info('deleted successfully');
        this.getTeamCount()
        this.getTeamLists(this.pageIndex, this.pageSize)
        this.loaderSvc.dismiss();
      },
      error: (error) => {
        this.loaderSvc.dismiss();
        this.apiSvc.handleErrorType(error);
      }
    })
  }

  startEdit(id: string) {
    this.clickedId = id;
    const selectedData = this.teamData.find(data => data._id === id);
    this.teamEditForm.patchValue({
      name: selectedData?.name,
      department_id: selectedData?.department_id,
    });
  }

  cancelEdit(): void {
    this.clickedId = '';
    this.messageSvc.info('cancelled');
  }

  updateDepartmentNames(): void {
    if (this.teamData && this.departmentData) {
      this.teamData.forEach((team: Team) => {
        const department = this.departmentData.find((dept: Department) => dept._id === team.department_id);
        if (department) {
          team.department_name = department.name;
        }
      });
    }
  }

  pageIndexChanged(pageIndex: number): void {
    this.pageIndex = pageIndex;
    this.getTeamLists(pageIndex, this.pageSize);
  }

  updateTeam(id: string) {
    this.loaderSvc.call();
    this.teamSvc.updateTeam(id, this.teamEditForm.value).subscribe({
      next: () => {
        this.messageSvc.info(`team ${this.teamEditForm.value.name} updated successfully`)
        this.getTeamLists(this.pageIndex, this.pageSize)
        this.clickedId = '';
        this.loaderSvc.dismiss();
      },
      error: (error) => {
        this.loaderSvc.dismiss();
        this.apiSvc.handleErrorType(error);
      }
    })
  }

  createUserModal(): void {
    const modalRef = this.modalSvc.create({
      nzContent: TeamCreateComponent,
      nzFooter: null,
      nzClosable: true,
      nzCentered: true,
      nzStyle: { 'width': '550px' },
      nzNoAnimation: true,
      nzMaskClosable: false,
      nzComponentParams: {
        departmentData: this.departmentData
      }
    });
    if (this.teamData) {
      modalRef?.componentInstance?.teamCreated.subscribe((teamData: Team) => {
        this.teamData = [...this.teamData, teamData]
        this.updateDepartmentNames();
        this.getTeamCount();
      });
    }
  }
}
