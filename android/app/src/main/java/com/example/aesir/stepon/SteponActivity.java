package com.example.aesir.stepon;

import android.os.Bundle;
import android.support.design.widget.FloatingActionButton;
import android.support.design.widget.Snackbar;
import android.support.v7.app.AppCompatActivity;
import android.support.v7.widget.Toolbar;
import android.view.View;
import android.view.Menu;
import android.view.MenuItem;
import android.widget.Button;
import android.widget.EditText;
import android.widget.TextView;

import go.stepon.Stepon;

public class SteponActivity extends AppCompatActivity {

    private EditText mUnits;
    private EditText mPrice;
    private TextView mTotal;
    private Button mCalc;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_stepon);
        Toolbar toolbar = (Toolbar) findViewById(R.id.toolbar);
        setSupportActionBar(toolbar);

        mUnits = (EditText) findViewById(R.id.mUnits);
        mPrice = (EditText) findViewById(R.id.mPrice);
        mCalc = (Button) findViewById(R.id.mCalc);
        mTotal = (TextView) findViewById(R.id.mTotal);

        mCalc.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                long units = Long.parseLong(mUnits.getText().toString());
                double price = Double.parseDouble(mPrice.getText().toString());
                String state = "UT";
                mTotal.setText(Stepon.GetTotal(units, price, state).toString());
            }
        });
    }
}
