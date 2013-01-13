<%@ Page Title="" Language="C#" MasterPageFile="~/Site.master" AutoEventWireup="true" CodeFile="register.aspx.cs" Inherits="_Default" %>

<asp:Content ID="Content1" ContentPlaceHolderID="HeadContent" Runat="Server">
    <style type="text/css">
        .style2
        {
            width: 100%;
        }
        .style3
        {
            text-align: right;
        }
    </style>
</asp:Content>
<asp:Content ID="Content2" ContentPlaceHolderID="MainContent" Runat="Server">
    <div style="height: 459px">
        <br />
        <br />
        
        <br />
        <br />
        <table class="style2">
            <tr>
                <td class="style3">
                    Registration</td>
                <td>
                    &nbsp;</td>
            </tr>
            <tr>
                <td class="style3">
                    &nbsp;</td>
                <td>
                    &nbsp;</td>
            </tr>
            <tr>
                <td class="style3">
                    User name:</td>
                <td>
                    <asp:TextBox ID="TextBox1" runat="server" Width="168px"></asp:TextBox>
                </td>
            </tr>
            <tr>
                <td class="style3">
                    Password:</td>
                <td>
                    <asp:TextBox ID="TextBox2" runat="server" Width="168px"></asp:TextBox>
                </td>
            </tr>
            <tr>
                <td class="style3">
                    Project:</td>
                <td>
                    <asp:DropDownList ID="DropDownList1" runat="server" Height="16px" Width="213px">
                    </asp:DropDownList>
                </td>
            </tr>
            <tr>
                <td class="style3">
                    City:</td>
                <td>
                    <asp:TextBox ID="TextBox3" runat="server" Width="168px"></asp:TextBox>
                </td>
            </tr>
            <tr>
                <td class="style3">
                    State:</td>
                <td>
                    <asp:TextBox ID="TextBox4" runat="server" Width="168px"></asp:TextBox>
                </td>
            </tr>
            <tr>
                <td class="style3">
                    Country</td>
                <td>
                    <asp:TextBox ID="TextBox5" runat="server" Width="168px"></asp:TextBox>
                </td>
            </tr>
            <tr>
                <td class="style3">
                    System config:Processor - RAM (Optional):</td>
                <td>
                    <asp:TextBox ID="TextBox6" runat="server" Width="168px"></asp:TextBox>
                </td>
            </tr>
            <tr>
                <td class="style3">
                    Internet Speed available:</td>
                <td>
                    <asp:TextBox ID="TextBox7" runat="server" Width="168px"></asp:TextBox>
                </td>
            </tr>
            <tr>
                <td class="style3">
                    &nbsp;</td>
                <td>
                    &nbsp;</td>
            </tr>
            <tr>
                <td class="style3">
                    &nbsp;</td>
                <td>
                    &nbsp;</td>
            </tr>
            <tr>
                <td class="style3">
                    &nbsp;</td>
                <td>
                    <asp:Button ID="Button1" runat="server" onclick="Button1_Click" Text="Register" 
                        Width="134px" />
                </td>
            </tr>
        </table>
        
    </div>

</asp:Content>

