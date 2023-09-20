import { Component, EventEmitter, Input, Output } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { NzModalRef, NzModalService } from 'ng-zorro-antd/modal';
import { TeamService } from 'src/app/repositories/team.service';
import { NzMessageService } from "ng-zorro-antd/message"
import { Department } from 'src/app/interfaces/department';
import { Team } from 'src/app/interfaces/team';
import { LoaderService } from 'src/app/services/loader.service';
import { ApiService } from 'src/app/services/api.service';

@Component({
  selector: 'app-team-create',
  templateUrl: './team-create.component.html',
  styleUrls: ['./team-create.component.scss']
})
export class TeamCreateComponent {
  @Input() departmentData: Department[] = [];
  @Output() teamCreated: EventEmitter<any> = new EventEmitter<any>();
  teamCreateForm!: FormGroup;
  isConfirmLoading = false;
  constructor(
    private modal: NzModalRef,
    private formBuilder: FormBuilder,
    private teamSvc: TeamService,
    private messageSvc: NzMessageService,
    private modalnew: NzModalService,
    private loaderSvc: LoaderService,
    private apiSvc: ApiService
  ) {
    this.teamCreateForm = this.formBuilder.group({
      name: ['', Validators.required],
      department_id: ['', Validators.required]
    });
  }

  cancelCreate() {
    this.modal.close();
  }

  createTeam() {
    this.isConfirmLoading = true;
    this.loaderSvc.call();
    this.teamSvc.createTeam(this.teamCreateForm.value).subscribe({
      next: (teamData: Team) => {
        this.messageSvc.info(`team "${this.teamCreateForm.value.name}" created successfully`);
        this.teamCreated.emit(teamData);
        this.isConfirmLoading = false;
        this.modal.close();
        this.loaderSvc.dismiss();
      },
      error: (error) => {
        if (error.status === 401) {
          this.loaderSvc.dissAll();
        }
        this.loaderSvc.dismiss();
        this.isConfirmLoading = false;
        this.apiSvc.handleErrorType(error);
      }
    })
  }

  showPopconfirm(): void {
    this.modalnew.confirm({
      nzTitle: 'Are you sure you want to create this team?',
      nzOkText: 'OK',
      nzCancelText: 'Cancel',
      nzCentered: true,
      nzOnOk: () => {
        this.createTeam(),
          this.isConfirmLoading = false;
      },
      nzOnCancel: () => {
        this.modal.destroy();
      },
    });
  }
}
