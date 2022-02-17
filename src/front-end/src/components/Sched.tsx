import {
    Inject, ScheduleComponent,
    Day, Week, WorkWeek, Month, Agenda,
    EventSettingsModel,
    ViewsDirective,
    ViewDirective
} from '@syncfusion/ej2-react-schedule';

export const Sched = () => {
    return (
    <ScheduleComponent
        currentView='Week'
        readonly={true}
    >
        <ViewsDirective>
            <ViewDirective
                option='Week'
                startHour='6:00'
                endHour='24:00'
            />
        </ViewsDirective>

        <Inject services={[Day, Week, WorkWeek, Month, Agenda]} />
    </ScheduleComponent>
    );
}

