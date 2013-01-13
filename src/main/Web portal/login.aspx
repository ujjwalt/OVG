<%@ Page Title="" Language="C#" MasterPageFile="~/Site.master" AutoEventWireup="true" CodeFile="login.aspx.cs" Inherits="_Default" %>

<asp:Content ID="Content1" ContentPlaceHolderID="HeadContent" Runat="Server">
    <style type="text/css">
        .style2
        {
            width: 100%;
            height: 72px;
        }
        .style3
        {
            text-align: right;
        }
    </style>
</asp:Content>
<asp:Content ID="Content2" ContentPlaceHolderID="MainContent" Runat="Server">
    <p>
        <br />
    </p>
    <div style="height: 248px">
        <table class="style2">
            <tr>
                <td class="style3">
                    User Login</td>
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
                    Username:</td>
                <td>
                    <asp:TextBox ID="TextBox1" runat="server" Width="184px"></asp:TextBox>
                </td>
            </tr>
            <tr>
                <td class="style3">
                    Password:</td>
                <td>
                    <asp:TextBox ID="TextBox2" runat="server" Width="184px" TextMode="Password"></asp:TextBox>
                </td>
            </tr>
            <tr>
                <td class="style3">
                    &nbsp;</td>
                <td>
                    <asp:Label ID="Label1" runat="server" Text="Enter correct values"></asp:Label>
                </td>
            </tr>
            <tr>
                <td class="style3">
                    &nbsp;</td>
                <td>
                    <asp:Button ID="Button1" runat="server" onclick="Button1_Click" Text="Login" 
                        Width="114px" />
                </td>
            </tr>
        </table>
    </div>
</asp:Content>

