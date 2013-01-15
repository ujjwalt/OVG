using System;
using System.Collections.Generic;
using System.Linq;
using System.Web;
using System.Web.UI;
using System.Web.UI.WebControls;
using System.Data;
using System.Data.SqlClient;

public partial class _Default : System.Web.UI.Page
{
    protected void Page_Load(object sender, EventArgs e)
    {

    }
    protected void Button1_Click(object sender, EventArgs e)
    {

        SqlConnection con = new SqlConnection(@"Data Source=.\SQLEXPRESS;AttachDbFilename=C:\Users\Technoworld\Documents\Test.mdf;Integrated Security=True;Connect Timeout=30;User Instance=True");
        con.Open();
        SqlDataReader read;

        Random rand = new Random((int)DateTime.Now.Ticks);
        int numIterations = 0;
        numIterations = rand.Next(1000, 10000);

        string s = "";
        s = "insert into Register values('" + numIterations + "','" + TextBox1.Text + "','" + TextBox2.Text + "','" + DropDownList1.SelectedItem + "','" + TextBox3.Text + "','" + TextBox4.Text + "','" + TextBox5.Text + "','" + TextBox6.Text + "','" + TextBox7.Text + "');";

        SqlCommand cmd = new SqlCommand(s, con);
        read = cmd.ExecuteReader();

        Response.Redirect("~/login.aspx", true);

    }
}