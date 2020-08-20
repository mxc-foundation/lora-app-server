import classNames from 'classnames';
import React, {Component} from "react";
import {Link} from "react-router-dom";
import {Breadcrumb, BreadcrumbItem, Card, Col, Row} from 'reactstrap';
import AdvancedTable from "../../components/AdvancedTable";
import Loader from "../../components/Loader";
import TitleBar from "../../components/TitleBar";
import TitleBarButton from "../../components/TitleBarButton";
import i18n, {packageNS} from '../../i18n';
import OrganizationStore from "../../stores/OrganizationStore";
import {MAX_DATA_LIMIT} from '../../util/pagination';


class ListOrganizations extends Component {
    constructor(props) {
        super(props);

        this.handleTableChange = this.handleTableChange.bind(this);
        this.getPage = this.getPage.bind(this);

        this.getColumns = this.getColumns.bind(this);
        this.organizationNameColumn = this.organizationNameColumn.bind(this);
        this.canHaveGatewaysColumn = this.canHaveGatewaysColumn.bind(this);

        this.state = {
            data: [],
            totalSize: 0,
            loading: false,
        }
    }

    organizationNameColumn = (cell, row, index, extraData) => {
        return <Link to={`/organizations/${row.id}`}>{row.name}</Link>;
    };

    canHaveGatewaysColumn = (cell, row, index, extraData) => {
        return <i className={classNames("mdi", {
            "mdi-check": row.canHaveGateways,
            "mdi-close": !row.canHaveGateways
        }, "font-20")}></i>;
    };

    serviceProfileColumn = (cell, row, index, extraData) => {
        return <div>
            <div>
                <Link to={`/organizations/${row.id}/service-profiles/create`}>ADD</Link>
            </div>
            <div>
                <Link to={`/organizations/${row.id}/service-profiles`}>CHECK</Link>
            </div>
        </div>;
    };

    getColumns = () => (
        [
            {
                dataField: 'name',
                text: i18n.t(`${packageNS}:tr000042`),
                sort: false,
                formatter: this.organizationNameColumn,
            },
            {
                dataField: 'displayName',
                text: i18n.t(`${packageNS}:tr000126`),
                sort: false,
            },
            {
                dataField: 'canHaveGateways',
                text: i18n.t(`${packageNS}:tr000380`),
                sort: false,
                formatter: this.canHaveGatewaysColumn,
            },
            {
                dataField: 'serviceProfiles',
                text: i18n.t(`${packageNS}:tr000078`),
                sort: false,
                formatter: this.serviceProfileColumn,
            }
        ]
    );

    /**
     * Handles table changes including pagination, sorting, etc
     */
    handleTableChange = (type, {page, sizePerPage, filters, sortField, sortOrder}) => {
        const offset = (page - 1) * sizePerPage;

        this.getPage(sizePerPage, offset);
    };

    getPage = async (limit, offset) => {
        limit = MAX_DATA_LIMIT;
        this.setState({loading: true});

        const res = await OrganizationStore.list("", limit = 10, offset = 0);
        //console.log(res)
        const object = this.state;
        object.totalSize = Number(res.totalCount);
        object.data = res.result;
        object.loading = false;
        this.setState({object});
    }

    componentDidMount() {
        this.getPage(MAX_DATA_LIMIT, 0);
    }

    render() {

        return (
            <React.Fragment>
                <TitleBar buttons={
                    <TitleBarButton
                        key={1}
                        label={i18n.t(`${packageNS}:tr000277`)}
                        icon={<i className="mdi mdi-plus mr-1 align-middle"></i>}
                        to={`/organizations/create`}
                    />}
                >
                    <Breadcrumb>
                        <BreadcrumbItem>{i18n.t(`${packageNS}:menu.control_panel`)}</BreadcrumbItem>
                        <BreadcrumbItem active>{i18n.t(`${packageNS}:tr000049`)}</BreadcrumbItem>
                    </Breadcrumb>
                </TitleBar>

                <Row>
                    <Col>
                        <Card className="card-box shadow-sm">
                            {this.state.loading && <Loader/>}
                            <AdvancedTable data={this.state.data} columns={this.getColumns()} keyField="id"
                                           totalSize={this.state.totalSize}
                                           onTableChange={this.handleTableChange}></AdvancedTable>
                        </Card>
                    </Col>
                </Row>
            </React.Fragment>
        );
    }
}

export default ListOrganizations;
