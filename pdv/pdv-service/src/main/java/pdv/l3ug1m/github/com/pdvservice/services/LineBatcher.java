package pdv.l3ug1m.github.com.pdvservice.services;

import java.io.Closeable;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

public class LineBatcher implements Closeable {
	
	private List<String[]> lines;
	private int size;
	private FlusingHandler flusingHandler;
	
	public  LineBatcher(int size, FlusingHandler flusingHandler) {
		this.size = size;
		this.flusingHandler = flusingHandler;
		this.lines = new ArrayList<>(size);
	}
	
	private void initValues() {
		this.lines = new ArrayList<String[]>(this.size);
	}

	public void add(String[] line) {
		this.lines.add(line);
		if (this.lines.size() == this.size) {
			this.flush();
		}
	}

	public void flush() {
		this.flusingHandler.handler(this.lines);
		initValues();
	}
	
	public static interface FlusingHandler {
		void handler(List<String[]> lines);
	}

	@Override
	public void close() throws IOException {
		this.flush();
	}

}
