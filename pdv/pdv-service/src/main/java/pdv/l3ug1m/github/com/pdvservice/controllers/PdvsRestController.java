package pdv.l3ug1m.github.com.pdvservice.controllers;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import pdv.l3ug1m.github.com.pdvservice.services.PdvService;
import pdv.l3ug1m.github.com.pdvservice.services.PdvService.PdvBean;

@RestController
public class PdvsRestController {

	private static final org.slf4j.Logger log = org.slf4j.LoggerFactory.getLogger(PdvsRestController.class);

	@Autowired
	PdvService pdvService;

	@GetMapping
		public List<PdvBean> list(@RequestParam(value = "tenant", required = true) String tenantId,
				@RequestParam(value = "nome", required = false) String nome,
				@RequestParam(value = "cidade", required = false) String cidade,
				@RequestParam(value = "endereco", required = false) String endereco,
				@RequestParam(value = "cep", required = false) String cep,
				@RequestParam(value = "offset", required = false) Integer offset) {

		return this.pdvService.list(tenantId, nome, cidade, endereco, cep, offset);
	}

	@ExceptionHandler
	public ResponseEntity<?> handleStorageFileNotFound(Throwable exc) {
		log.error(exc.getMessage(), exc);
		return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).build();
	}

}
