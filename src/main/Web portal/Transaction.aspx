<%@ Page Title="" Language="C#" MasterPageFile="~/Site.master" AutoEventWireup="true" CodeFile="Transaction.aspx.cs" Inherits="_Default" %>

<asp:Content ID="Content1" ContentPlaceHolderID="HeadContent" Runat="Server">
</asp:Content>
<asp:Content ID="Content2" ContentPlaceHolderID="MainContent" Runat="Server">

<div style="height: 280px">

    <br />
    <br />
    <h1>Make Transaction</h1>
    <p>
        <asp:GridView ID="GridView1" runat="server" AutoGenerateColumns="False" 
            CellPadding="4" DataSourceID="SqlDataSource1" ForeColor="#333333" 
            GridLines="None" Width="630px">
            <AlternatingRowStyle BackColor="White" />
            <Columns>
                <asp:BoundField DataField="clientID" HeaderText="clientID" 
                    SortExpression="clientID" />
                <asp:BoundField DataField="projectID" HeaderText="projectID" 
                    SortExpression="projectID" />
                <asp:BoundField DataField="credits" HeaderText="credits" 
                    SortExpression="credits" />
                <asp:BoundField DataField="balanace" HeaderText="balanace" 
                    SortExpression="balanace" />
                <asp:ButtonField Text="Transfer Money" />
            </Columns>
            <EditRowStyle BackColor="#2461BF" />
            <FooterStyle BackColor="#507CD1" Font-Bold="True" ForeColor="White" />
            <HeaderStyle BackColor="#507CD1" Font-Bold="True" ForeColor="White" />
            <PagerStyle BackColor="#2461BF" ForeColor="White" HorizontalAlign="Center" />
            <RowStyle BackColor="#EFF3FB" />
            <SelectedRowStyle BackColor="#D1DDF1" Font-Bold="True" ForeColor="#333333" />
            <SortedAscendingCellStyle BackColor="#F5F7FB" />
            <SortedAscendingHeaderStyle BackColor="#6D95E1" />
            <SortedDescendingCellStyle BackColor="#E9EBEF" />
            <SortedDescendingHeaderStyle BackColor="#4870BE" />
        </asp:GridView>
        <asp:SqlDataSource ID="SqlDataSource1" runat="server" 
            ConnectionString="<%$ ConnectionStrings:TestConnectionString %>" 
            SelectCommand="SELECT * FROM [credits]"></asp:SqlDataSource>
    </p></div>
</asp:Content>

