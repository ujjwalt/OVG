using System;
using System.Collections.Generic;
using System.Linq;
using System.Web;
using System.Web.UI;
using System.Web.UI.WebControls;
using System.IO;
using System.Data.SqlClient;
using System.Data;

public partial class _Default : System.Web.UI.Page
{
    protected void Page_Load(object sender, EventArgs e)
    {

        
    }
    protected void Button1_Click(object sender, EventArgs e)
    {
        
        Label1.Visible = false;
        SqlConnection con = new SqlConnection(@"Data Source=.\SQLEXPRESS;AttachDbFilename=C:\Users\Technoworld\Documents\Test.mdf;Integrated Security=True;Connect Timeout=30;User Instance=True");
        con.Open();
        SqlDataReader read;
        string cm="select password from Register where username= '" + TextBox1.Text + "';";
        
            SqlCommand cmd=new SqlCommand(cm,con);
        read=cmd.ExecuteReader();
        
        //Response.Write("Hi");
        //if (read.Read())
        if (read.HasRows)
        {

            //Response.Write(read["password"]);
            //if (read["password"].ToString() == TextBox2.Text)
            // {
            Response.Redirect("~/projects.aspx");
            // }
        }
        else
        {

            Label1.Visible = true;
        }

        
      //  else
        //    Label1.Visible = true;


    }
}