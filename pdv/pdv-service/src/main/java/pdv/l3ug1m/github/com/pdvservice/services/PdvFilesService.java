package pdv.l3ug1m.github.com.pdvservice.services;

import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.nio.file.StandardCopyOption;
import java.util.LinkedList;
import java.util.List;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.Future;

import org.jooq.DSLContext;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.web.multipart.MultipartFile;

import com.opencsv.CSVReader;
import com.opencsv.CSVReaderBuilder;

import pdv.l3ug1m.github.com.pdvservice.controllers.StorageException;
import pdv.l3ug1m.github.com.pdvservice.jooq.tables.Pdvs;
import pdv.l3ug1m.github.com.pdvservice.jooq.tables.records.PdvsRecord;
import pdv.l3ug1m.github.com.pdvservice.services.LineBatcher.FlusingHandler;

@Service
public class PdvFilesService {

	@Autowired
	DSLContext dslContext;

	@Autowired
	PdvService pdvService;

	Path rootLocation = Paths.get("/pdv-service/update-dir");
//	Path rootLocation = Paths.get("update-dir");

	public Path store(MultipartFile file, String tenantId) {
		String filename = tenantId+"-"+System.currentTimeMillis();
		try {
			if (file.isEmpty()) {
				throw new StorageException("Failed to store empty file " + filename);
			}
			if (filename.contains("..")) {
				// This is a security check
				throw new StorageException(
						"Cannot store file with relative path outside current directory " + filename);
			}
			Path path = this.rootLocation.resolve(filename);
			try (java.io.InputStream inputStream = file.getInputStream()) {
				Files.copy(inputStream, path, StandardCopyOption.REPLACE_EXISTING);
			}
			return path;
		} catch (IOException e) {
			throw new StorageException("Failed to store file " + filename, e);
		}
	}

	@Transactional
	public void teste() {
		PdvsRecord newRecord = this.dslContext.newRecord(Pdvs.PDVS);

		newRecord.setTenant("tenant");
		newRecord.setNome("nome");
		newRecord.setEndereco("endereco");
		newRecord.setCep("cep");

		newRecord.insert();
	}

	public void processFile(Path pathFile, final String tenantId) {
		final List<Future<Boolean>> futureResult = new LinkedList<>();
		FlusingHandler handler = new FlusingHandler() {

			@Override
			public void handler(List<String[]> lines) {
				futureResult.add(pdvService.processLines(tenantId, lines));
			}
		};

		try (CSVReader csvReader = new CSVReaderBuilder(new BufferedReader(new FileReader(pathFile.toFile())))
				.withSkipLines(1)
				.build();
				LineBatcher lineBatcher = new LineBatcher(3000, handler)) {
			String[] line = null;
			while ((line = csvReader.readNext()) != null) {
				lineBatcher.add(line);
			}

		} catch (IOException e) {
			throw new RuntimeException(e.getMessage(), e);
		}

		try {
			for (Future<Boolean> future : futureResult) {
				future.get();
			}
		} catch (InterruptedException | ExecutionException e) {
			throw new RuntimeException(e.getMessage(), e);
		}

	}

}
